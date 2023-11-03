import React, {useEffect, useState} from "react";
import {base64} from "../utils/base64";

function Register() {
     async function registration() {
        const optionsResponse = await fetch("https://f84e-99-228-180-110.ngrok-free.app/api/v1/auth/registerRequest");
        const {publicKey} = await optionsResponse.json()

        // Base64URL decode some values
        publicKey.user.id = base64.decode(publicKey.user.id);
        publicKey.challenge = base64.decode(publicKey.challenge);
        if (publicKey.excludeCredentials) {
            for (let cred of publicKey.excludeCredentials) {
                cred.id = base64.decode(cred.id);
            }
        }

        // Use platform authenticator and discoverable credential
        publicKey.authenticatorSelection = {
            authenticatorAttachment: 'platform',
            requireResidentKey: true
        }

        // Invoke WebAuthn create
        const cred = await navigator.credentials.create({
            publicKey: publicKey,
        }) as PublicKeyCredential;

        if (cred === null) return;

        // Base64URL encode some values
        const clientDataJSON = base64.encode(cred.response.clientDataJSON);
        const attestationObject = base64.encode((cred.response as AuthenticatorAttestationResponse).attestationObject);

        // Obtain transports if they are available.
        const transports = (cred.response as AuthenticatorAttestationResponse).getTransports ? (cred.response as AuthenticatorAttestationResponse).getTransports() : [];

        const credential = {
            id: cred.id,
            rawId: base64.encode(cred.rawId),
            type: cred.type,
            authenticatorAttachment: cred.authenticatorAttachment ? cred.authenticatorAttachment : undefined,
            response: {
                clientDataJSON,
                attestationObject,
                transports
            }
        };

        const headers = {
            "X-Requested-With": "XMLHttpRequest",
            "Content-Type": "application/json",
        };

        // Send the result to the server and return the promise.
        return await fetch("https://f84e-99-228-180-110.ngrok-free.app/api/v1/auth/registerResponse",
            {method: "POST", body: JSON.stringify(credential), credentials: "same-origin", headers: headers});
    }

    return (
        <div className="Register">
            <button onClick={registration}>Register</button>
        </div>
    );
}

export default Register;
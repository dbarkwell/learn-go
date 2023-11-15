import React, {useState} from "react";
import {base64} from "../utils/base64";

function Register() {
    const [ message, setMessage ] = useState("");
    async function registration(event: React.FormEvent<EventTarget>) {
        event.preventDefault();
        const target = event.target as HTMLFormElement;
        const username = (target.elements.namedItem("username") as HTMLInputElement).value;
        const headers = {
             "X-Requested-With": "XMLHttpRequest",
             "Content-Type": "application/json",
        };
        const optionsResponse = await fetch(`/api/v1/auth/registerRequest/${username}`,
            {method: "POST", credentials: "same-origin", headers: headers});

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



        // Send the result to the server and return the promise.
        const registerResponse = await fetch(`/api/v1/auth/registerResponse/${username}`,
            {method: "POST", body: JSON.stringify(credential), credentials: "same-origin", headers: headers});
        const { message } = await registerResponse.json();
        setMessage(message);
    }

    return (
        <div className="Register">
            <form onSubmit={registration}>
                <div>Username</div>
                <div><input type={"text"} name={"username"} /></div>
                <button type={"submit"} value={"Register"}>Register</button>
            </form>
            <div>{message}</div>
        </div>
    );
}

export default Register;
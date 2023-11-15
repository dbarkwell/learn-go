import {base64} from "../utils/base64";
import React, {useState} from "react";

function Login() {
    const [ message, setMessage ] = useState("");
    async function login(event: React.FormEvent<EventTarget>) {
        event.preventDefault();
        const target = event.target as HTMLFormElement;
        const username = (target.elements.namedItem("username") as HTMLInputElement).value;
        const optionsResponse = await fetch(`/api/v1/auth/signinRequest/${username}`);
        const {publicKey} = await optionsResponse.json();
        publicKey.challenge = base64.decode(publicKey.challenge);

        // `allowCredentials` empty array invokes an account selector by discoverable credentials.
        publicKey.allowCredentials = [];

        // Invoke WebAuthn get
        const cred = await navigator.credentials.get({
            publicKey: publicKey,
            // Request a conditional UI
            mediation: 'optional'
        }) as PublicKeyCredential;

        if (!cred) return;

        // Base64URL encode some values
        const clientDataJSON = base64.encode(cred.response.clientDataJSON);
        const authenticatorData = base64.encode((cred.response as AuthenticatorAssertionResponse).authenticatorData);
        const signature = base64.encode((cred.response as AuthenticatorAssertionResponse).signature);
        const userHandle = base64.encode((cred.response as AuthenticatorAssertionResponse).userHandle);

        const credential = {
            id: cred.id,
            type: cred.type,
            rawId: base64.encode(cred.rawId),
            response: {
                clientDataJSON,
                authenticatorData,
                signature,
                userHandle,
            }
        };

        const headers = {
            "X-Requested-With": "XMLHttpRequest",
            "Content-Type": "application/json",
        };

        const signInResponse = await fetch(`/api/v1/auth/signinResponse/${username}`,
            {method: "POST", body: JSON.stringify(credential), credentials: "same-origin", headers: headers});
        const { message } = await signInResponse.json();
        setMessage(message);
    }

    return (
        <div className="Login">
            <form onSubmit={login}>
                <div>Username</div>
                <div><input type={"text"} name={"username"} /></div>
                <button type={"submit"} value={"Login"}>Login</button>
            </form>
            <div>{message}</div>
        </div>
    );
}

export default Login;
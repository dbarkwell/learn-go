import {base64} from "../utils/base64";

function Login() {
    async function signIn() {
        const optionsResponse = await fetch("https://f84e-99-228-180-110.ngrok-free.app/api/v1/auth/signinRequest");
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

        await fetch("https://f84e-99-228-180-110.ngrok-free.app/api/v1/auth/signinResponse",
            {method: "POST", body: JSON.stringify(credential), credentials: "same-origin", headers: headers});
    }

    return (
        <div className="Login">
            <button onClick={signIn}>Sign In</button>
        </div>
    );
}

export default Login;
import React from "react";

interface NewUser {
    username: string
    firstName: string
    lastName: string
    email: string
}

function Users() {

    async function submitForm(event: React.FormEvent<EventTarget>) {
        event.preventDefault();
        const target = event.target as HTMLFormElement;

        const newUser: NewUser = {
            username: (target.elements.namedItem("username") as HTMLInputElement).value,
            firstName: (target.elements.namedItem("firstName") as HTMLInputElement).value,
            lastName: (target.elements.namedItem("lastName") as HTMLInputElement).value,
            email: (target.elements.namedItem("email") as HTMLInputElement).value,
        };

        const headers = {
            "X-Requested-With": "XMLHttpRequest",
            "Content-Type": "application/json",
        };

        await fetch("/api/v1/users",
            {method: "POST", body: JSON.stringify(newUser), credentials: "same-origin", headers: headers});
    }

    return (
        <div className="Users">
            <form onSubmit={submitForm}>
                <div>Username</div>
                <div><input type={"text"} name={"username"} /></div>
                <div>First Name</div>
                <div><input type={"text"} name={"firstName"} /></div>
                <div>Last Name</div>
                <div><input type={"text"} name={"lastName"} /></div>
                <div>Email</div>
                <div><input type={"text"} name={"email"} /></div>
                <button type={"submit"} value={"Save"}>Save</button>
            </form>
        </div>
    );
}

export default Users;
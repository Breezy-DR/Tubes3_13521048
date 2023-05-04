import { useState } from "react";
import axios from "axios";

function CreateUser() {
    const [name, setName] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        const response = await axios.post("http://localhost:3000/users", { name });
        console.log(response.data);
        setName("");
    };

    return (
        <div>
            <h1>Create User</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    Name:
                    <input type="text" value={name} onChange={(e) => setName(e.target.value)} />
                </label>
                <button type="submit">Submit</button>
            </form>
        </div>
    );
}
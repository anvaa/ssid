
async function delClick() {
        
    var messageElement = document.getElementById("_message");
    messageElement.style.border = "1px solid red";

    var uid = document.getElementById("_uid").value;
    var email = document.getElementById("_email").value;

    const verify = confirm("Are you sure you want to delete user " + email + 
                            "? \n\nYou can just remove auth to deny access.", "Delete user");
    if (!verify) {
        return;
    }

    var userData = {
        id: uid,
    };

    try {
        const response = await fetch("/user/delete/" + uid, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(userData),
        });
    
        if (!response.ok) {
            messageElement.innerHTML = error.message;
        }
    
        window.location.href = "/v/user/" + uid;
        } catch (error) {
        messageElement.innerHTML = "Delete failed: " + error.message;

    }
}

async function setPswClick() {
        
    var messageElement = document.getElementById("_message");
    
    var uid = document.getElementById("_uid").value;
    var psw1 = document.getElementById("_password1").value;
    var psw2 = document.getElementById("_password2").value;

    if (!validatePasswords(psw1, psw2)) {
        return; // Message is set inside the validatePasswords function
    }

    var userData = {
        id: uid,
        password: psw1,
    };

    try {
        const response = await fetch("/user/psw", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(userData),
        });
    
        if (!response.ok) {
            messageElement.innerHTML = error.message;
            messageElement.style.border = "1px solid red";
        }
    
        window.location.href = "/v/user/" + uid;
        } catch (error) {
        messageElement.innerHTML = "Change password failed: " + error.message;
        messageElement.style.border = "1px solid red";
    }
}

function validatePasswords(password, password2) {
    var messageElement = document.getElementById("_message");
    messageElement.style.border = "1px solid red";

    if (password === "" || password2 === "") {
        messageElement.innerHTML = "Passwords cannot be empty";
    return false;
    }

    if (password !== password2) {
        messageElement.innerHTML = "Passwords do not match";
    return false;
    }

    if (password.length < 8) {
    messageElement.innerHTML = "Passwords must be at least 8 characters long";
    return false;
    }

    return true;
}

async function setRole() {
        
    var messageElement = document.getElementById("_message");

    var id = document.getElementById("_uid").value;
    var role = document.getElementById("_role").value;
    
    var userData = {
        id: id,
        role: role,
    };

    try {
        const response = await fetch("/user/role", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(userData),
        });
    
        if (!response.ok) {
            messageElement.innerHTML = error.message;
            messageElement.style.border = "1px solid red";
        } 

        window.location.href = "/v/user/" + id;
        } catch (error) {
        messageElement.innerHTML = "Change role failed: " + error.message;
        messageElement.style.border = "1px solid red";
    }
}

async function setAccessTime() {
        
    var messageElement = document.getElementById("_message");
    
    var uid = document.getElementById("_uid").value;
    var min = document.getElementById("_min").value;

    var userData = {
        id: uid,
        accesstime: min,
    };

    try {
        const response = await fetch("/user/act", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(userData),
        });
    
        if (!response.ok) {
            messageElement.innerHTML = error.message;
            messageElement.style.border = "1px solid red";
        }
    
        window.location.href = "/v/user/" + uid;
        } catch (error) {
        messageElement.innerHTML = "Change access time failed: " + error.message;
        messageElement.style.border = "1px solid red";
    }
}

async function setUrl() {
        
    var messageElement = document.getElementById("_message");

    var id = document.getElementById("_uid").value;
    var url = document.getElementById("_url").value;
    
    var userData = {
        id: id,
        url: url,
    };

    try {
        const response = await fetch("/user/url", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(userData),
        });
    
        if (!response.ok) {
            messageElement.innerHTML = error.message;
            messageElement.style.border = "1px solid red";
        } 

        window.location.href = "/v/user/" + id;
    } catch (error) {
        messageElement.innerHTML = "Change url failed: " + error.message;
        messageElement.style.border = "1px solid red";
    }
}
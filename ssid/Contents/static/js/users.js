
async function setAuth(id) {
    
    const isAuth = document.getElementById("_auth" + id).value;
    const messageElement = document.getElementById("message");
    const returl = window.location.pathname;

    var data = {
        id: id,
        isauth: isAuth,
    };
    
    try {
        const response = await fetch("/user/auth", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            messageElement.innerHTML = error.message;
            messageElement.style.border = "1px solid red";
            return;
        }

        window.location.href = returl;
    } catch (error) {
        messageElement.innerHTML = "Change auth failed: " + error.message;
        messageElement.style.border = "1px solid red";
    }
}
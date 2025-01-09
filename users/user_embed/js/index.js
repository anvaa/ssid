
document.addEventListener("keydown", function (event) {
  if (event.key === "Enter") {
  
    if (document.getElementById("_pageid").value == "login") {
      loginClick();
    }

    if (document.getElementById("_pageid").value == "signup") {
      signupClick();
    }

  }
});


async function loginClick() {

  const email = document.getElementById("_email").value;
  const password = document.getElementById("_password").value;

if (!validEmail(email)) {
    writeMessage("Invalid email");
    return;
  }

  if (!validPassword(password)) {
    writeMessage("Invalid password");
    return;
  }

  const data = {
    email: email,
    password: password,
  };

  try {
    const response = await fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data),
    });

    const responseData = await response.json();
    
    if (response.status === 301) {
      window.location.href = responseData.url;
    } else {
      throw new Error(responseData.message);
    }
    
  } catch (error) { 
    writeMessage("Login: " + error.message);
  }
};

async function signupClick() {
    
  const email = document.getElementById("_email").value;
  const password = document.getElementById("_password").value;
  const password2 = document.getElementById("_password2").value;
  
  if (!validatePasswords(password, password2)) {
    return; // Message is set inside the validatePasswords function
  }
  
  try {
    const userData = {email, password, password2 };
    
    const response = await fetch("/signup", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(userData),
    });

    const responseData = await response.json();

    if (!response.ok) {
      throw new Error('Server error');
    }

    window.location.href = responseData.url;
  } catch (error) {
    writeMessage("Signup: " + error.message);
  }
};
    
function validatePasswords(psw1, psw2) {
  
  if (psw1 === "" || psw2 === "") {
    writeMessage("Passwords cannot be empty");
    return false;
  }

  if (psw1 !== psw2) {
    writeMessage("Passwords do not match");
    return false;
  }

  if (psw1.length < 8) {
    writeMessage("Passwords must be at least 8 characters long");
    return false;
  }

  if (psw1.length > 50) {
    writeMessage("Passwords must be less than 50 characters long");
    return false;
  }

  return true;
}
  

function validPassword(psw) {
  return psw.length >= 8;
}

function validEmail(email) {
  return email.length >= 8;
}

function writeMessage(message) {
  const messageElement = document.getElementById("_message");
  messageElement.innerHTML = message;
  messageElement.style.borderColor = "red";
}

document.addEventListener("keydown", function (event) {
    if (event.key === "Escape") {
        resetPage();
    }

    if (event.key == "+") {
        itmAddNewClick();
    }

    if (event.key === "F1") {
        Home();
    }
    
    if (event.key === "F2") {
        Search();
    }

    if (event.key === "F3") {
        Reports();
    }

    if (event.key === "F4") {
        Tools();
    }

    if (event.key === "F12") {
        Logout();
    }
});

function Home() {
    window.location.href = "/app";
}

function Search() {
    window.location.href = "/app/search";
}

function Stats() {
    window.location.href = "/app/stats";
}

function Tools() {
    window.location.href = "/app/tools";
}

function Info() {
    window.location.href = "/info";
}

function Logout() {
    window.location.href = "/logout";
}

function itmAddNewClick() {
    
    var url = "/itm/new";
    var xhr = new XMLHttpRequest();
    xhr.open("GET", url, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
            window.location.href = "/itm/new";
           
        }
    }
    xhr.send();
}
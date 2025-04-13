
document.addEventListener("keydown", function (event) {
    if (event.key === "Escape") {
        resetPage();
    }
});

document.addEventListener("keydown", function (event) {
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
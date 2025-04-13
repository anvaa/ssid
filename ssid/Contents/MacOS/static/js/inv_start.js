
function resetPage() {
    window.location.reload();
    return;
};

// listen for keydown event
document.addEventListener("keydown", function (event) {
    if (event.key == "+") {
        itmAddNewClick();
    }
});

document.addEventListener("keydown", function (event) {
    if (event.key == "Enter") {
        if (document.getElementById("_siteaction").value = "SearchId") {
            var itmid = document.getElementById("_txtSearchId").value;
            if (itmid != "") {
                itmQuickSearch('ID',itmid);
            }
        }
        if (document.getElementById("_siteaction").value = "SearchSerial") {
            var snr = document.getElementById("_txtSearchSerial").value;
            if (snr != "") {
                itmQuickSearch('SNR',snr);
            }
        }
    }
});

function setSearch(searchitm) {
    document.getElementById("_siteaction").value = searchitm; 
}

function invHomeClick() {
    
    var url = "/app/home";
    var xhr = new XMLHttpRequest();
    xhr.open("GET", url, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
            document.getElementById("div_home").innerHTML = xhr.responseText;
            document.getElementById("_siteaction").value = "home";
            document.getElementById("_txtSearchSerial").focus();
        }
    }
    xhr.send();
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

async function itmQuickSearch(act,search_val) {

    var url = "";
    if (act == "ID") {
        // if search_val is numeric
        if (!isNaN(search_val)) {
            url = "/search/"+search_val;
        } else {
            alert("Please enter a valid ID");
            return;
        }
    } 

    if (act == "SNR") {
        url = "/search/serial/"+search_val;
    }
    
    // redirect to search page
    window.location.href = url;
    
}

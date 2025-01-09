


function resetPage() {
    window.location.reload();
    return;
};

function setFocus() {
    document.getElementById("_locsel").focus();
}

async function itmSearch(srcval) {

    var locsel = document.getElementById("_locsel").value;
    var typsel = document.getElementById("_typsel").value;
    var mansel = document.getElementById("_mansel").value;
    var stasel = document.getElementById("_stasel").value;

    var body = {
        "locid": parseInt(locsel),
        "typid": parseInt(typsel),
        "manid": parseInt(mansel),
        "staid": parseInt(stasel),
    };

    if (locsel== 0 && typsel == 0 && mansel == 0 && stasel == 0) {
        alert("Please select search criterias");
        return;
    }

    if (srcval == "PRINT") {
        exSearch(body);
        return;
    }
    
    var url = "/search/multi";
    var xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
            document.getElementById("div_searchresult").innerHTML = xhr.responseText;
            prntpage = xhr.responseText;
            
        } 
    }
    xhr.send(JSON.stringify(body));
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

async function exSearch(data) {
    // use search data to make a csv file to download

    var divdata = document.getElementById("div_searchresult").innerHTML;
    
    if (divdata.length < 760) {
        alert("Nothing to download");
        return;
    }

    try {
        const response = await fetch("/search/export", {
            method: "POST",
            headers: { "Content-Type": "application/file" },
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            throw new Error("Server error");
        }

        const responseData = await response.blob();
        const url = window.URL.createObjectURL(responseData);

        // get filename from response header
        const contentDisposition = response.headers.get("content-disposition");
        var filename = contentDisposition.split(";")[1].split("filename=")[1].trim();

        filename = removeFirstAndLastChar(filename);

        // download file using a tmp tag
        const a = document.createElement("a");
        a.href = url;
        a.download = filename;
        document.body.appendChild(a);
        a.click();
        window.URL.revokeObjectURL(url);
        
    } catch (error) {
        alert("Error: " + error.message);
    }
}

function removeFirstAndLastChar(str) {
    return str.length > 2 ? str.slice(1, -1) : '';
}
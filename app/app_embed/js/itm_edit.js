
function resetPage() {
    window.location.reload();
    return;
};

document.addEventListener("keydown", function (event) {
    if (event.key === "Enter") {
        if (document.getElementById("_siteaction").value == "edit_itm") {
            itmAddUpdate();
        }
    }
});

function isEdit() {
    document.getElementById("_isedit").value = "true";
}

function setFocus() {
    document.getElementById("_locsel").focus();
}

async function itmAddUpdate() {

    if (document.getElementById("_isedit").value == "false") {
        alert("Nothing to add/update");
        return;
    }

    
    var locsel = document.getElementById("_locsel").value;
    var typsel = document.getElementById("_typsel").value;
    var mansel = document.getElementById("_mansel").value;
    var stasel = document.getElementById("_stasel").value;
    
    var itmid = document.getElementById("_itmid0").value;
    var itmuid = document.getElementById("_itmuid0").value;
    var itmdesc = document.getElementById("_itmdesc0").value;
    var itmserial = document.getElementById("_itmserial0").value;
    var itmprice = document.getElementById("_itmprice0").value;
    
    if (locsel == "" || typsel == "" || mansel == "") {
        alert("Please select item location, type and manufacturer");
        return;
    }

    if (itmdesc == "" || itmserial == "" || itmprice == "") {
        alert("Please fill in text fields");
        return;
    }

    // replace
    itmprice = itmprice.replace(",", ".");
    itmprice = itmprice.trim();

    // var url = window.location.pathname;
    var data = {
        "locid":  parseInt(locsel),
        "typid":  parseInt(typsel),
        "manid":  parseInt(mansel),
        "staid":  parseInt(stasel),
       
        "Description": itmdesc,
        "serial": itmserial,
        "price": parseFloat(itmprice) ,
        "uid": parseInt(itmuid),
        "itmid": parseInt(itmid),
    };

    try {
        const response = await fetch("/itm/addupd", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data),
        });

        const responseData = await response.json();
        if (!response.ok) {
            throw new Error(responseData.Error);
        }

        window.location.href = "/app";
    } catch (error) {
        alert("Add Item Status failed: " + error.message);
    }
}

function itmDelete(itmid) {
    
    if (!confirm("Are you sure you want to delete this item?")) {
        return;
    }

    var url = "/itm/delete/"+itmid;
    var xhr = new XMLHttpRequest();
    xhr.open("DELETE", url, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
            window.location.href = "/app";
        }
    }
    xhr.send();
}

function itmPrintQrClick() {
    // print image
    var itmid = document.getElementById("_itmid0").value;
    var img = document.getElementById("_qrimg").src;
    var win = window.open("", "_blank");
    
    printpage = "<div style='width: 100%;text-align: center;'>"+
                    "<img src='"+img+"' style='width: 100%;margin-bottom: 0px;' />"+
                    "<div style='font-family: monospace;font-size: 48px;font-weight: bold;margin-top: 0px;'>"+itmid+"</div>"+
                "</div>";
               
    win.document.write(printpage);
    
    win.print();
    win.close();
   
}

async function staAddStatusHist(itmid) {
    var uid = document.getElementById("_uid0").value;
    var stasel = document.getElementById("_stasel").value;
    var stacom = document.getElementById("_stacom0").value;
    
    if (stasel == "" ) {
        alert("Please select a status");
        return;
    }
    
    if (stacom == "") {
        alert("Please enter a comment");
        return;
    }
    
    var url = "/sta/hist/add";
    var data = {
        "itmid": parseInt(itmid),
        "uid": parseInt(uid),
        "staid" : parseInt(stasel),
        "txt" : stacom
    };
    
    try {
        const response = await fetch(url, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data),
        });

        const responseData = await response.json();
        if (!response.ok) {
            throw new Error(responseData.Error);
        }

        window.location.reload();
    } catch (error) {
       alert("Error: " + error);
    }

}
document.addEventListener("keydown", handleKeyDown);

function handleKeyDown(event) {
    
    if (event.key === "Enter" && document.getElementById("_siteaction").value === "edit_itm") {
        itmAddUpdate();
    }
}

function resetPage() {
    window.location.reload();
}

function isEdit() {
    document.getElementById("_isedit").value = "true";
}

function setFocus() {
    document.getElementById("_locsel").focus();
}

async function itmAddUpdate() {
    if (document.getElementById("_isedit").value === "false") {
        alert("Nothing to add/update");
        return;
    }

    const locsel = document.getElementById("_locsel").value;
    const typsel = document.getElementById("_typsel").value;
    const mansel = document.getElementById("_mansel").value;
    const stasel = document.getElementById("_stasel").value;
    const itmid = document.getElementById("_itmid0").value;
    const itmuid = document.getElementById("_itmuid0").value;
    const itmdesc = document.getElementById("_itmdesc0").value;
    const itmserial = document.getElementById("_itmserial0").value;
    let itmprice = document.getElementById("_itmprice0").value;

    if (!locsel || !typsel || !mansel) {
        alert("Please select item location, type and manufacturer");
        return;
    }

    if (!itmdesc || !itmserial || !itmprice) {
        alert("Please fill in text fields");
        return;
    }

    itmprice = itmprice.replace(",", ".").trim();

    const data = {
        locid: parseInt(locsel),
        typid: parseInt(typsel),
        manid: parseInt(mansel),
        staid: parseInt(stasel),
        Description: itmdesc,
        serial: itmserial,
        price: parseFloat(itmprice),
        uid: parseInt(itmuid),
        itmid: parseInt(itmid),
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

    const url = `/itm/delete/${itmid}`;
    const xhr = new XMLHttpRequest();
    xhr.open("DELETE", url, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            window.location.href = "/app";
        }
    };
    xhr.send();
}

function imgPrint(prtype) {
    const itmid = document.getElementById("_itmid0").value;
    const img = document.getElementById(prtype).src;
    const win = window.open("", "_blank");
    
    const printpage = `
        <div style='width: 100%; text-align: center;'>
            <img src='${img}' style='width: 100%; margin-bottom: 0px;' />
            <div style='font-family: monospace; font-size: 48px; font-weight: bold; margin-top: 0px;'>${itmid}</div>
        </div>
    `;

    win.document.write(printpage);
    win.print();
    win.close();
}

async function staAddStatusHist(itmid) {
    const uid = document.getElementById("_uid0").value;
    const stasel = document.getElementById("_stasel").value;
    const stacom = document.getElementById("_stacom0").value;

    if (!stasel) {
        alert("Please select a status");
        return;
    }

    if (!stacom) {
        alert("Please enter a comment");
        return;
    }

    const data = {
        itmid: parseInt(itmid),
        uid: parseInt(uid),
        staid: parseInt(stasel),
        txt: stacom,
    };

    try {
        const response = await fetch("/sta/hist/add", {
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
        alert("Error: " + error.message);
    }
}

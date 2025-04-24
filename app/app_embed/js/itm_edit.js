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
    const stasel = document.getElementById("_stasel0").value;
    const itmid = document.getElementById("_itmid0").value;
    const itmuid = document.getElementById("_itmuid0").value;
    let itmdesc = document.getElementById("_itmdesc0").value;
    let itmserial = document.getElementById("_itmserial0").value;
    let itmprice = document.getElementById("_itmprice0").value;

    if (!locsel || !typsel || !mansel) {
        alert("Please select item location, type and manufacturer");
        return;
    }

    if (itmdesc === "") {
        itmdesc = "Nil";
    }

    if (itmserial === "") {
        itmserial = "Nil";
    }

    if (itmprice === "") {
        itmprice = "0.0";
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

async function addStatus() {

    const itmid = parseInt(document.getElementById("_itmid").value);
    const staid = parseInt(document.getElementById("_stasel0").value);
    const uid = parseInt(document.getElementById("_uid0").value);
    let txt = document.getElementById("_comm0").value;

    if (!staid) {
        alert("Please select a status");
        return;
    }

    if (txt === "") {
        txt = "Nil";
    }

    const data = { itmid, uid, staid, txt };

    try {
        const response = await fetch("/sta/hist/add", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            const responseData = await response.json();
            throw new Error(responseData.Error);
        }

        window.location.reload();
    } catch (error) {
        alert(`Error: ${error.message}`);
    }
}

async function delStatus(stahist_id) {

    const staid = document.getElementById("_staid0").value;

    if (staid === "1212090603") {
        alert("Cannot delete this status");
        return;
    }

    if (!confirm("Are you sure you want to delete this status ?")) {
        return;
    }
    
    const url = `/sta/delete/${stahist_id}`;
    const xhr = new XMLHttpRequest();
    xhr.open("DELETE", url, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            window.location.reload();
        }
    };
    xhr.send();

}

function imgPrint(imgtype) {

    const serial = document.getElementById("_itmserial0").value;
    const img = document.getElementById(imgtype).src;
    const win = window.open("", "_blank");

    const imgConfig = {
        _barimg: { height: '15mm', width: '100%', margin: '2mm', page_width: '1mm', page_height: '0mm', text_align: 'center' },
        _qrimg: { height: '16mm', width:'16mm', margin: '2mm', page_width: '2mm', page_height: '0mm', text_align: 'left' },
    };

    const txtConfig = {
        _barimg: `<div style="width: inherit;">${serial}</div>`,
        _qrimg: "",
    };

    const img_height = imgConfig[imgtype].height;
    const img_width = imgConfig[imgtype].width;
    const img_margin = imgConfig[imgtype].margin;
    const img_align = imgConfig[imgtype].text_align;
    const img_txt = txtConfig[imgtype];

    if (!img_height || !img_width || !img_margin) {
        alert("Invalid image type");
        return;
    }

    const printLabel = `
        <html>
            <head>
                <title>Print</title>
                <style>
                    @media print {
                        body {
                            margin: 0;
                            padding: 0;
                            margin: ${img_margin};
                        }
                        img {
                            width: ${img_width};
                            height: ${img_height};
                            padding: 3px;
                        }
                        @page {
                            size: ${imgConfig[imgtype].page_width} ${imgConfig[imgtype].page_height};
                        }
                        .print-label {
                            width: 100%;
                            height: 100%;
                            margin: 0;
                            text-align: ${img_align};
                            font-size: 12px;
                            font-family: monospace;
                            font-weight: bold;
                        }
                    }
                </style>
            </head>
            <body>
                <div class="print-label">
                    <img src="${img}" />
                    ${img_txt}
                </div>
            </body>
        </html>`;

    win.document.write(printLabel);
    win.document.close();
    win.print();
    win.close();
}

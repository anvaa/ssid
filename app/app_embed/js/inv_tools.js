
var isedit;
var selTxt;
var selId;

function resetPage() {
    window.location.reload();
    return;
};

document.addEventListener("keydown", function (event) {
    if (event.key === "Enter") {
        
        if (isedit == "") {
            return;
        }
        selTxt = document.getElementById("_"+isedit+"txt0").value;
        lstAddUpd();
    }
});

function isEdit(val) {
    isedit = val;
}

// LOCATIONS //
function locSel(id) {
    isedit = 'loc';
    selId = id;
    selTxt = document.getElementById("_selloc"+selId).value;
    document.getElementById("_loctxt0").value = selTxt;
    document.getElementById("_locid0").value = id;
} 

// TYPES //
function typSel(id) {
    isedit = 'typ';
    selId = id;
    selTxt = document.getElementById("_seltyp"+selId).value;
    document.getElementById("_typtxt0").value = selTxt;
    document.getElementById("_typid0").value = id;
} 

// MANUFACT //
function manSel(id) {
    isedit = 'man';
    selId = id;
    selTxt = document.getElementById("_selman"+selId).value;
    document.getElementById("_mantxt0").value = selTxt;
    document.getElementById("_manid0").value = id;
} 

// STATUS //
function staSel(id) {
    isedit = "sta";
    selId = id;
    selTxt = document.getElementById("_selsta"+selId).value;
    document.getElementById("_statxt0").value = selTxt;
    document.getElementById("_staid0").value = id;
} 


async function lstAddUpd() {
    var txt = document.getElementById("_"+isedit+"txt0").value;
    var id = document.getElementById("_"+isedit+"id0").value;

    if (txt == "") {
        alert("Status: Nothing to add or update");
        return; 
    }

    if (id == "") {
        id = 0;
    }

    var data = {
        "txt": txt,
        "id": id,
        "url": window.location.pathname,
    };
    
    try {
        const response = await fetch("/"+isedit+"/addupd", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data),
        });

        const responseData = await response.json();
        if (!response.ok) {
            throw new Error(responseData.Error);
        }
        
        window.location.href = responseData.url;
    } catch (error) {
        alert("Add/Update "+isedit+" failed: " + error.message);
    }
}

async function lstDel() {

    var txt = document.getElementById("_"+isedit+"txt0").value;

    if (txt == "") {
        alert("Nothing to delete");
        return; 
    }

    if (selId == "1212090603") {
        alert("Can't delete status 'New'");
        return;
    }

    if (!confirm("Delete '"+txt+"'?")) {
        return;
    }

    var data = {
        "id": selId,
        "url": window.location.pathname,
    };
    
    try {
        const response = await fetch("/"+isedit+"/delete", {
            method: "DELETE",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data),
        });

        const responseData = await response.json(); 
       if (response.status != 200) {
            throw new Error(responseData.error);
        }

        window.location.href = responseData.url;
    } catch (error) {
        alert("Delete "+isedit+" failed: " + error.message);
    }
}
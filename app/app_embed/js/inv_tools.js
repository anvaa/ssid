
function resetPage() {
    window.location.reload();
    return;
};

document.addEventListener("keydown", function (event) {
    if (event.key === "Enter") {
        var iseditval = document.getElementById("_isedit").value;
        if (iseditval == "") {
            return;
        }
        if (iseditval == "loc") {
            locAddUpd();
        }
        if (iseditval == "typ") {
            typAddUpd();
        }
        if (iseditval == "man") {
            manAddUpd();
        }
        if (iseditval == "sta") {
            staAddUpd();
        }
    }
});

function isEdit(editval) {
    document.getElementById("_isedit").value = editval;
}

// LOCATIONS //
function locSel(id) {
    
    document.getElementById("_loctxt0").value = document.getElementById("_selloc"+id).value;
    document.getElementById("_locid0").value = id;
} 

// TYPES //
function typSel(id) {
    
    document.getElementById("_typtxt0").value = document.getElementById("_seltyp"+id).value;
    document.getElementById("_typid0").value = id;
} 

// MANUFACT //
function manSel(id) {
    
    document.getElementById("_mantxt0").value = document.getElementById("_selman"+id).value;
    document.getElementById("_manid0").value = id;
} 

// STATUS //
function staSel(id) {
    
    document.getElementById("_statxt0").value = document.getElementById("_selsta"+id).value;
    document.getElementById("_staid0").value = id;
} 


async function lstAddUpd(lstitm) {
    var txt = document.getElementById("_"+lstitm+"txt0").value;
    var id = document.getElementById("_"+lstitm+"id0").value;

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
        const response = await fetch("/"+lstitm+"/addupd", {
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
        alert("Add/Update "+lstitm+" failed: " + error.message);
    }
}

async function lstDel(lstitm) {

    var txt = document.getElementById("_"+lstitm+"txt0").value;
    var id = document.getElementById("_"+lstitm+"id0").value;

    if (txt == "") {
        alert("Nothing to delete");
        return; 
    }

    if (id == "1212090603") {
        alert("Can't delete status 'New'");
        return;
    }

    if (!confirm("Delete '"+txt+"'?")) {
        return;
    }

    var data = {
        "id": id,
        "url": window.location.pathname,
    };
    
    try {
        const response = await fetch("/"+lstitm+"/delete", {
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
        alert("Delete "+lstitm+" failed: " + error.message);
    }
}
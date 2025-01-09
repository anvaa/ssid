
async function staAddStatus(itmid) {
    var stat = document.getElementById("_stasel0").value;
    var txt = document.getElementById("_stacom0").value;
    var uid = document.getElementById("_uid0").value;
    alert("stat: " + stat + ", txt: " + txt + ", uid: " + uid);
    if (txt == "" || stat == "") {
        alert("Nothing to add!");
        return; 
    }

    var data = {
        id: itmid,
        stat: stat,
        txt: txt,
        uid: uid,
    };
    
    try {
        const response = await fetch("/sta/hist/add", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            throw new Error("Server error");
        }

        window.location.href = "/app";
    } catch (error) {
        alert("Add Item Status failed: " + error.message);
    }
}
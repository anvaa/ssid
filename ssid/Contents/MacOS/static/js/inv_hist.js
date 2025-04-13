async function staAddStatus(itmid) {
    const stat = document.getElementById("_stasel0").value;
    const txt = document.getElementById("_stacom0").value;
    const uid = document.getElementById("_uid0").value;

    if (!stat || !txt) {
        alert("Nothing to add!");
        return;
    }

    const data = { id: itmid, stat, txt, uid };

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

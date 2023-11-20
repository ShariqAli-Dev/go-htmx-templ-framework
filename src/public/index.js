// google oauth
const googleAuth = document.getElementById("google-auth");
googleAuth?.addEventListener("click", () => {
    window.location.href += "dashboard";
});
// increment state
const incrementButton = document.getElementById("incrementButton");
const decrementButton = document.getElementById("decrementButton");
const counter = document.getElementById("counter");
incrementButton?.addEventListener("click", () => {
    counter.innerText = (parseInt(counter.innerText) + 1).toString();
});
decrementButton?.addEventListener("click", () => {
    counter.innerText = (parseInt(counter.innerText) - 1).toString();
});
// create person
const createPersonButton = document.getElementById("createPersonButton");
const createdPersonPre = document.getElementById("createdPerson");
createPersonButton?.addEventListener("click", async () => {
    createPersonButton.disabled = true;
    const shariq = await createPerson({ name: "Shariq", id: 1 });
    createPersonButton.disabled = false;
    createdPersonPre.innerText = JSON.stringify(shariq, null, 2);
});
// dashboard modal
document.addEventListener("DOMContentLoaded", () => {
    const dashboardModal = document.getElementById("dashboard-modal");
    if (dashboardModal) {
        const dashboardModalCloseDate = localStorage.getItem("dashboardModalCloseDate");
        if (dashboardModalCloseDate) {
            const modalCloseDate = new Date(dashboardModalCloseDate);
            const today = new Date();
            const diff = today.getTime() - modalCloseDate.getTime();
            const diffInMinutes = diff / (1000 * 60);
            if (diffInMinutes >= 1) {
                dashboardModal.showModal();
            }
        }
        else {
            dashboardModal.showModal();
        }
    }
});
const dashboardModalCloseButton = document.getElementById("dashbord-modal-close");
dashboardModalCloseButton?.addEventListener("click", () => {
    localStorage.setItem("dashboardModalCloseDate", new Date().toString());
});
// functions
async function createPerson(person) {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve(person);
        }, 2000);
    });
}

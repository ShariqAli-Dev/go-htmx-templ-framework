// @ts-expect-error
lucide.createIcons();
// elements
const incrementButton = document.getElementById("incrementButton");
const decrementButton = document.getElementById("decrementButton");
const counter = document.getElementById("counter");
const createPersonButton = document.getElementById("createPersonButton");
const createdPersonPre = document.getElementById("createdPerson");
const dashboardModal = document.getElementById("dashboard-modal");
const dashboardModalCloseButton = document.getElementById("dashbord-modal-close");
const dashboardModalOpenButton = document.getElementById("open-dashboard-modal-button");
// increment state
incrementButton?.addEventListener("click", () => {
    counter.innerText = (parseInt(counter.innerText) + 1).toString();
});
decrementButton?.addEventListener("click", () => {
    counter.innerText = (parseInt(counter.innerText) - 1).toString();
});
// google oauth
const googleAuth = document.getElementById("google-auth");
googleAuth?.addEventListener("click", () => {
    window.location.href += "dashboard";
});
// create person
createPersonButton?.addEventListener("click", async () => {
    createPersonButton.disabled = true;
    const shariq = await createPerson({ name: "Shariq", id: 1 });
    createPersonButton.disabled = false;
    createdPersonPre.innerText = JSON.stringify(shariq, null, 2);
});
// dashboard modal
document.addEventListener("DOMContentLoaded", () => {
    if (dashboardModal) {
        const dashboardModalCloseDate = localStorage.getItem("dashboardModalCloseDate");
        if (dashboardModalCloseDate) {
            const modalCloseDate = new Date(dashboardModalCloseDate);
            const today = new Date();
            const diff = today.getTime() - modalCloseDate.getTime();
            const diffInDays = diff / (1000 * 60 * 60 * 24);
            if (diffInDays >= 1) {
                dashboardModal.showModal();
            }
        }
        else {
            dashboardModal.showModal();
        }
    }
});
dashboardModalCloseButton?.addEventListener("click", () => {
    localStorage.setItem("dashboardModalCloseDate", new Date().toString());
});
dashboardModalOpenButton?.addEventListener("click", () => {
    dashboardModal.showModal();
});
// functions
async function createPerson(person) {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve(person);
        }, 2000);
    });
}

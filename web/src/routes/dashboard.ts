// quiz me
const createQuizSection = document.getElementById(
  "create-quiz"
) as HTMLDivElement;
createQuizSection?.addEventListener("click", () => {
  window.location.href = window.location.origin + "/quiz";
});

// dashboard modal
const dashboardModal = document.getElementById(
  "dashboard-modal"
) as HTMLDialogElement;
const dashboardModalCloseButton = document.getElementById(
  "dashbord-modal-close"
) as HTMLButtonElement;
const dashboardModalOpenButton = document.getElementById(
  "open-dashboard-modal-button"
) as HTMLButtonElement;

document.addEventListener("DOMContentLoaded", () => {
  if (dashboardModal) {
    const dashboardModalCloseDate = localStorage.getItem(
      "dashboardModalCloseDate"
    );
    if (dashboardModalCloseDate) {
      const modalCloseDate = new Date(dashboardModalCloseDate);
      const today = new Date();
      const diff = today.getTime() - modalCloseDate.getTime();
      const diffInDays = diff / (1000 * 60 * 60 * 24);

      if (diffInDays >= 1) {
        dashboardModal.showModal();
      }
    } else {
      dashboardModal.showModal();
    }
  }
});

dashboardModalCloseButton?.addEventListener("click", () => {
  localStorage.setItem("dashboardModalCloseDate", new Date().toString());
});

dashboardModalOpenButton?.addEventListener("click", async () => {
  dashboardModal.showModal();
});

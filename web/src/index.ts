// ***** INIT **** //
// @ts-expect-error
lucide.createIcons();

// ***** CONSTANTS ***** //
const DASHBOARD_PATH = window.location.origin + "/dashboard";
const QUIZ_PATH = window.location.origin + "/quiz";

// ***** INDEX ***** //
const createQuizSection = document.getElementById(
  "create-quiz"
) as HTMLDivElement;
const title = document.getElementById("title") as HTMLHeadingElement;

//heading nav click
title?.addEventListener("click", () => {
  if (window.location.href !== DASHBOARD_PATH) {
    window.location.href = DASHBOARD_PATH;
  }
});
// google oauth
const googleAuth = document.getElementById("google-auth");
googleAuth?.addEventListener("click", () => {
  window.location.href += "dashboard";
});

// ***** DASHBOARD ***** //
// quiz me
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
// ***** QUIZ **** //
type Variant = "multiple-choice" | "open-ended";

const multipleChoiceButton = document.getElementById(
  "multiple-choice-button"
) as HTMLButtonElement;
const openEndedButton = document.getElementById(
  "open-ended-button"
) as HTMLButtonElement;
const variantInput = document.getElementById("variant-input");

// switching variant
multipleChoiceButton?.addEventListener("click", () => {
  if (variantInput?.getAttribute("value") !== ("multiple-choice" as Variant)) {
    openEndedButton.classList.remove("btn-neutral");
    multipleChoiceButton.classList.add("btn-neutral");
    variantInput?.setAttribute("value", "multiple-choice" as Variant);
  }
});
openEndedButton?.addEventListener("click", () => {
  if (variantInput?.getAttribute("value") !== ("open-ended" as Variant)) {
    multipleChoiceButton.classList.remove("btn-neutral");
    openEndedButton.classList.add("btn-neutral");
    variantInput?.setAttribute("value", "open-ended" as Variant);
  }
});

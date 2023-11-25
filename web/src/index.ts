// init
// @ts-expect-error
lucide.createIcons();

const DASHBOARD_PATH = window.location.origin + "/dashboard";

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

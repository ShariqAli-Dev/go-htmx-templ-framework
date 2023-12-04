// init
// @ts-expect-error
lucide.createIcons();
const INDEX_PATH = window.location.origin + "/";
// responsive nav
const signInAnchor = document.getElementById("sign-in") as HTMLAnchorElement;
const signOutAnchor = document.getElementById("sign-out") as HTMLAnchorElement;
if (
  window.location.href === INDEX_PATH ||
  window.location.href === window.location.origin
) {
  signInAnchor?.classList.remove("hidden");
  signOutAnchor?.classList.add("hidden");
} else {
  signInAnchor?.classList.add("hidden");
  signOutAnchor?.classList.remove("hidden");
}

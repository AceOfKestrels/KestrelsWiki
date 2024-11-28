function toggleSidebarLeft() {
    document.getElementById("__layout").classList.toggle("left-collapsed")
    document.getElementById("__toggle-left-button").classList.toggle("turned")
}

function toggleSidebarRight() {
    document.getElementById("__layout").classList.toggle("right-collapsed")
    document.getElementById("__toggle-right-button").classList.toggle("turned")
}
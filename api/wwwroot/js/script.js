function toggleSidebarLeft() {
    document.getElementById("layout").classList.toggle("left-collapsed")
    document.getElementById("toggle-left-button").classList.toggle("turned")
}

function toggleSidebarRight() {
    document.getElementById("layout").classList.toggle("right-collapsed")
    document.getElementById("toggle-right-button").classList.toggle("turned")
}
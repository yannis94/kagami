function remove() {
    document.querySelector(".popup-background").remove()
    document.querySelector("body").classList.remove("unscrollable")
}

document.querySelector(".popup-background").addEventListener("click", remove)
document.querySelector(".popup-cross").addEventListener("click", remove)

document.querySelector("body").classList.add("unscrollable")

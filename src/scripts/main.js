const nav = document.querySelector("nav ul")
const max = 300
const min = 50

const links = nav.querySelectorAll("a")
const currentPage = window.location.href

const animatedTitle = document.querySelector(".home-banner-title h1")
let titleContent = animatedTitle.textContent
let letters = []
let intervalInMilisecond = random()


links.forEach(link => {
    if (link.href === currentPage) {
        link.classList.add("current")
    }
})

animatedTitle.textContent = ""

for (const letter of titleContent) {
    letters.push(letter)
}

const block = document.createElement("div")
block.classList.add("title-block")
document.querySelector(".tapping").appendChild(block)

const animateInterval = setInterval(() => {
    animatedTitle.textContent += letters.splice(0, 1)
    if (letters.length === 0) {
        setTimeout(() => {
            block.remove()
        }, 1000)
        clearInterval(animateInterval)
    }
    intervalInMilisecond = random()
}, intervalInMilisecond)

function random() {
    return Math.floor(Math.random() * (max - min) + min)
}

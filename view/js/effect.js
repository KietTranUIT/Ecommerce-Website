window.addEventListener("scroll", function() {
    var announcebar = document.querySelector(".announcebar")
    var navigationbar = document.querySelector(".main-navigationbar")
    announcebar.classList.toggle("sticky", window.scrollY > 0)
    navigationbar.classList.toggle("test", window.scrollY > 0)
})

$(document).ready(function() {
    let submit_btn = $('.submit-button-register')
    let submit_svg = $('.submit-button-register svg')

    submit_btn.mouseenter(() => {
        submit_svg.css('fill', '#000000')
    })

    submit_btn.mouseleave(() => {
        submit_svg.css('fill', '#ffffff')
    })
})

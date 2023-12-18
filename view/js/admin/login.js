
$(document).ready(function() {
    // Gui yeu cau xac thuc dang nhap
    $('#btn-login').click(function() {
        var email = $('#email').val()
    var password = $('#password').val()

    let credential = {
        email: email,
        password: password
    }

    let dataJson = JSON.stringify(credential)
        console.log(dataJson)
        fetch("/admin/login", {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: dataJson
        })
        .then(response => {
            if (response.status === 200) {
                GotoAdminHome()
            } else {
                console.log("no")
            }
        })
    })
})

function GotoAdminHome() {
    fetch("/admin/products", {
        method: 'GET'
    })
    .then(response => {
        window.location.href = response.url
    })
}
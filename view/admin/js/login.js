
$(document).ready(function() {
    // Gui yeu cau xac thuc dang nhap
    $('#btn-submit').click(function(event) {
        event.preventDefault();
        var email = $('#email').val()
    var password = $('#password').val()

    let credential = {
        email: email,
        password: password
    }

    let dataJson = JSON.stringify(credential)
        fetch("/admin/login", {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: dataJson
        })
        .then(response => {
            if (response.status === 200) {
                alert("log in success!")
                GotoAdminHome()
            } else {
                alert("Log in failed!")
            }
        })
    })
})

function GotoAdminHome() {
    console.log("OK")
    fetch("/admin", {
        method: 'GET'
    })
    .then(response => {
        window.location.href = response.url
    })
}
$(document).ready(function() {
    $('form').submit(function(event) {
        event.preventDefault();

        let email = $('#email').val()
        let password = $('#password').val()

        let login_data = {
            email: email,
            password: password
        }

        console.log(login_data)

        fetch('/login', {
            method:'POST',
            headers: { 'Content-Type': 'application/js'},
            body: JSON.stringify(login_data)
        })
        .then(response => response.json())
        .then(response_data => {
            console.log(response_data)
        })
})
})
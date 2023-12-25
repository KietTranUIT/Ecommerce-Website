$(document).ready(function() {
    // Chuyen sang trang dang nhap
    $('#submit-register').click(function(event) {
        event.preventDefault()

        fetch('/signup', {
            method: 'GET'
        })
        .then(response => {
            if (response.status == 200) {
                window.location.href = response.url
            } else {
                alert("Error")
            }

        })
    })

    // Gui request den server de dang nhap
    $('#submit-login').click(function(event) {
        event.preventDefault();

        let email = $('#email').val()
        let password = $('#password').val()

        let login_data = {
            email: email,
            password: password
        }


        fetch('/login', {
            method:'POST',
            headers: { 'Content-Type': 'application/js'},
            body: JSON.stringify(login_data)
        })
        .then(response => response.json())
        .then(data => {
            if(!data.status) {
                alert("Username and password are incorrect!")
            } else {
                localStorage.setItem('user', data.data[0])
                window.location.href = '/'
            }
        })
})
})
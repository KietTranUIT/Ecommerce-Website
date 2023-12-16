
function IsEmail(email) {
    var regex = /^([a-zA-Z0-9_\.\-\+])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    if(!regex.test(email)) {
        return false;
    }else{
        return true;
    }
}

$(document).keypress(function(event) {
    if(event.which === 13) {
        event.preventDefault();
    }
})

$(document).ready(function() {
    $('#email').click(function() {
        $(this).css('border-color', '#dee2e6')
        $('#invalid-email').hide()
    })

    $('#email').keypress(function(event) {
        if(event.which === 13) {
            event.preventDefault()
            $('#password').focus()
        }
    })

    $('#password').keypress(function(event) {
        if(event.which === 13) {
            event.preventDefault()
        }
    })

    $('#password').click(function() {
        $(this).css('border-color', '#dee2e6')
        $('#invalid-password').hide()
    })

    $('form').submit(function(event) {
        event.preventDefault()
        Submit()
    })
})

async function Submit() {

        var buttonValue = $('form').find("button[type='submit']:focus").val();

        if(buttonValue === 'login') {
            let email = $('#email')
            let password = $('#password')
            console.log(email.val())

            if(email.val() === '' || !IsEmail(email.val())) {
                email.css('border-color', 'red')
                $('#invalid-email').show()
                return
            }

            if(password.val() === '') {
                password.css('border-color', 'red')
                $('#invalid-password').show()
                return  
            }

            let login_data = {
                email: email.val(),
                password: password.val()
            }

            await fetch('/login', {
                method: 'POST',
                headers: {'Content-Type': 'application/js'},
                body: JSON.stringify(login_data)
            })
            .then(response => response.json())
            .then(response_data => {
                if(!response_data['status']) {
                    if(response_data['error_code'] == 'LOGIN_FAIL') {
                        showAlert('#fail-login')
                    } else {
                        showAlert('#server-error')
                    }
                } else {
                    localStorage.setItem('user', JSON.stringify(response_data['data']))
                    window.location.href = '/'
                }
            })
        } else {
            window.location.href = '/signup'
        }
}

function showAlert(name) {
    let alert = $(name)
    alert.show()
    setInterval(function() {
        alert.hide()
    }, 5000)
}
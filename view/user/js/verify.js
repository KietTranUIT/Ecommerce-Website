function SendVerificationEmail(user) {
    let email = user.email
    fetch('/verify?email=' + email, {method: 'POST'})
    .then(response => response.json())
    .then(response_data => {
        if(response_data['error_code'] === "INTERNAL_ERROR") {
            showAlert('#server-error')
        } else if(response_data['error_code'] = 'DUPLICATE_ENTRY') {
            showAlert('#code-duplicate')
        } else {
            alert('Send verification email successfully')
            showAlert('#sendcode-success')
        }
    })
}

function showAlert(name) {
    let alert = $(name)
    alert.show()
    setInterval(function() {
        alert.hide()
    }, 5000)
}

$(document).ready(function() {
    let user = JSON.parse(localStorage.getItem('user-signup'))

    $('#formEmail').val(user.email)

    $('#code-number-1').on('keypress', function(event) {
        var charCode = event.which ? event.which : event.keyCode;

        // Kiểm tra nếu ký tự không phải là chữ số (0-9)
        if (charCode < 48 || charCode > 57) {
            event.preventDefault(); // Ngăn chặn hành vi mặc định nếu ký tự không phải là chữ số
        } else {
            $('#code-number-2').focus();
        }
    })

    $('#code-number-2').on('keypress', function(event) {
        var charCode = event.which ? event.which : event.keyCode;

        // Kiểm tra nếu ký tự không phải là chữ số (0-9)
        if (charCode < 48 || charCode > 57) {
            event.preventDefault(); // Ngăn chặn hành vi mặc định nếu ký tự không phải là chữ số
        } else {
            $('#code-number-3').focus();
        }
    })

    $('#code-number-3').on('keypress', function(event) {
        var charCode = event.which ? event.which : event.keyCode;

        // Kiểm tra nếu ký tự không phải là chữ số (0-9)
        if (charCode < 48 || charCode > 57) {
            event.preventDefault(); // Ngăn chặn hành vi mặc định nếu ký tự không phải là chữ số
        } else {
            $('#code-number-4').focus();
        }
    })

    $('#code-number-4').on('keypress', function(event) {
        var charCode = event.which ? event.which : event.keyCode;

        // Kiểm tra nếu ký tự không phải là chữ số (0-9)
        if (charCode < 48 || charCode > 57) {
            event.preventDefault(); // Ngăn chặn hành vi mặc định nếu ký tự không phải là chữ số
        } else {
            $('#code-number-5').focus();
        }
    })

    $('#code-number-5').on('keypress', function(event) {
        var charCode = event.which ? event.which : event.keyCode;

        // Kiểm tra nếu ký tự không phải là chữ số (0-9)
        if (charCode < 48 || charCode > 57) {
            event.preventDefault(); // Ngăn chặn hành vi mặc định nếu ký tự không phải là chữ số
        } else {
            $('#code-number-6').focus();
        }
    })

    $('#verify').click(function(event) {
        event.preventDefault()
        SignUp(user)
    })

    $('#resend').click(function(event) {
        console.log("resend")
        event.preventDefault()
        SendVerificationEmail(user)
    })
})

function SignUp(user) {
    var buttonValue = $('form').find("button[type='submit']:focus").val();

    if(buttonValue === 'resend') {
        return
    }

    let code = $('#code-number-1').val() + $('#code-number-2').val() + $('#code-number-3').val() + $('#code-number-4').val() + $('#code-number-5').val() + $('#code-number-6').val()

    let data_signup = {
        first_name: user.first_name,
        last_name: user.last_name,
        email: user.email,
        password: user.password,
        code: code,
        gender: user.gender,
        phone: user.phone,
        address: user.address,
        city: user.city
    }

    fetch('/signup', {
        method: 'POST',
        headers: {"Content-Type":"application/js"},
        body: JSON.stringify(data_signup)
    })
    .then(response => response.json())
    .then(response_data => {
        if(response_data['status']) {
            alert("Successfully registered account!")
            localStorage.removeItem('user-signup')
            window.location.href = '/login'
        } else {
            if(response_data['error_code'] == 'AUTHENTICATE_FAIL') {
                showAlert('#fail-verification')
            } else {
                showAlert('#server-error')
            }
        }
    })
}
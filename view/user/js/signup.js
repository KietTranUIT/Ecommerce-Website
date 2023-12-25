$(document).ready(function() {
    $('#create_customer').submit(function(event) {
        event.preventDefault()
    })

    $('#first_name').keypress(function(event) {
        if(event.which === 13) {
            event.preventDefault()
            $('#last_name').focus()
        } else {
            $(this).css('border-color', '#dee2e6')
            $('#error-first-name').hide()
        }
    })

    $('#last_name').keypress(function(event) {
        if(event.which === 13) {
            event.preventDefault()
            $('#email').focus()
        } else {
            $(this).css('border-color', '#dee2e6')
            $('#error-last-name').hide()
        }
    })

    $('#email').keypress(function(event) {
        if(event.which === 13) {
            event.preventDefault()
            $('#gender').focus()
        } else {
            $(this).css('border-color', '#dee2e6')
            $('#error-email').hide()
        }
    })

    $('#gender').keypress(function(event) {
        if(event.which === 13) {
            event.preventDefault()
            $('#phone_number').focus()
        } else {
            $(this).css('border-color', '#dee2e6')
            $('#error-gender').hide()
        }
    })

    $('#phone_number').keypress(function(event) {
        if(event.which === 13) {
            event.preventDefault()
            $('#address').focus()
        } else {
            $(this).css('border-color', '#dee2e6')
            $('#error-phone').hide()
        }
    })

    $('#address').keypress(function(event) {
        if(event.which === 13) {
            event.preventDefault()
            $('#password').focus()
        } else {
            $(this).css('border-color', '#dee2e6')
            $('#error-address').hide()
        }
    })

    $('#password').keypress(function(event) {
        if(event.which === 13) {
            event.preventDefault()
        } else {
            $(this).css('border-color', '#dee2e6')
            $('#error-').hide()
        }
    })
})


async function SignUp() {
        var first_name = $('#first_name')
        var last_name = $('#last_name')
        let email = $('#email')
        var password = $('#password')
        let gender = $('#gender')
        let phone_number = $('#phone_number')
        let address = $('#address')
        let city = $('#city')

        let flag = true
        if(first_name.val() === '') {
            flag = false
            first_name.css('border-color', 'red')
            $('#error-first-name').show()
        } 
        if(last_name.val() === '') {
            flag = false
            last_name.css('border-color', 'red')
            $('#error-last-name').show()
        }
        if(email.val() === '' || !IsEmail(email.val())) {
            flag = false
            email.css('border-color', 'red')
            $('#error-email').show()
        }
        if(phone_number.val() === '' || !IsPhonenumber(phone_number.val())) {
            flag = false
            phone_number.css('border-color', 'red')
            $('#error-phone').show()
        }
        if(address.val() === '') {
            flag = false
            address.css('border-color', 'red')
            $('#error-address').show()
        }
        if(password.val() === '') {
            flag = false
            password.css('border-color', 'red')
            $('#error-password').show()
        }
        if(city.val() === '') {
            flag = false
            city.css('border-color', 'red')
            $('#error-city').show()
        }

        if(!flag) {
            return
        } 

        let user = {
            first_name: first_name.val(),
            last_name: last_name.val(),
            email: email.val(),
            password: password.val(),
            gender: gender.val(),
            phone: phone_number.val(),
            address: address.val(),
            city: city.val()
        }
        localStorage.setItem("user-signup", JSON.stringify(user))

        let check = await isAccountExist(email.val())

        if(check) {
            let alert_email_already = $('#email-already')
            alert_email_already.show()
            setTimeout(function() {
                alert_email_already.hide()
            }, 5000)
            return
        }
        
        fetch('/verify', {
            method: 'GET'
        })
        .then(response => {
            window.location.href = response.url
            SendVerificationEmail(email.val())
        })
}

function IsEmail(email) {
    var regex = /^([a-zA-Z0-9_\.\-\+])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    if(!regex.test(email)) {
        return false;
    }else{
        return true;
    }
}

function IsPhonenumber(phonenumber) {
    var regex = /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im;
    if(!regex.test(phonenumber)) {
        return false;
    }else{
        return true;
    }
}

function isAccountExist(email) {
    return fetch('/account/check?email=' + email, {method: 'GET'})
    .then(response => response.json())
    .then(response_data => {
        if(response_data['error_code'] === 'DUPLICATE ENTRY') {
            return true
        }
        return false
    })
}

function SendVerificationEmail(email) {
    return fetch('/verify?email=' + email, {method: 'POST'})
}
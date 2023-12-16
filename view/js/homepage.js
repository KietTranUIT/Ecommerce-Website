$(document).ready(function() {
    let user_data = JSON.parse(localStorage.getItem('user'))

    if(user_data != null) {
        console.log(user_data)
        console.log(user_data[0].First_name + ' ' + user_data[0].Last_name)
        $('#user').text(user_data[0].First_name + ' ' + user_data[0].Last_name)
    }
})
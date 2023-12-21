
$(document).ready(function() {
    $('.log-out').click(function() {
        fetch('/admin/logout', {
            method: 'DELETE'
        })
        .then(response => {
            if(response.status === 200) {
                alert("Log out success")
                fetch('/admin/login', {
                    method: 'GET'
                })
                .then(response => {
                    if (response.status === 200) {
                        window.location.href = response.url
                    }
                })
            } else {
                alert("Log out fail")
            }
        })
    })

    // truy cap danh muc san pham
    $('#categories').click(function() {
        fetch('/admin/categories', {
            method: 'GET'
        })
        .then(response => {
            window.location.href = response.url
        })
    })

    // truy cap danh sach san pham
    $('#products').click(function() {
        fetch('/admin/products', {
            method: 'GET'
        })
        .then(response => {
            window.location.href = response.url
        })
    })

    // truy cap danh sach don hang
    $('#view-orders').click(function() {
        fetch('/admin/orders', {
            methos: 'GET'
        })
        .then(response => {
            window.location.href = response.url
        })
    })

    // truy cap trang chu
    $('#dashboard').click(function() {
        fetch('/admin', {
            methos: 'GET'
        })
        .then(response => {
            window.location.href = response.url
        })
    })
    $('.btn-detail').click(function() {
        let id = this.dataset.orderId
    
        let url = '/admin/orders/' + id
    
        fetch(url, {
            method: 'GET'
        })
        .then(response => {
                window.location.href = response.url
        })
    })
})
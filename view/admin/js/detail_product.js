$(document).ready(function() {
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
            methos: 'GET'
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
    $('.btn-delete').click(function() {
    var userConfirmed = window.confirm("Deleting a product will delete all data in the inventory")
    if (!userConfirmed) {
        return
    }
    let id = this.dataset.productId

    let url = '/admin/products_version/' + id

    fetch(url, {
        method: 'DELETE'
    })
    .then(response => {
            if (response.status === 200) {
                location.reload()
            }
        })
    })

    $('.btn-edit').click(function() {
        let id = this.dataset.productId
    
        let url = '/admin/products_version/update/' + id
    
        fetch(url, {
            method: 'GET'
        })
        .then(response => {
                window.location.href = response.url
            })
    })

    $('#create-new').click(function() {
        let id = this.dataset.productId
        let url = '/admin/products/' + id + '/new'
        fetch(url, {
            method: 'GET'
        })
        .then(response => {
            window.location.href = response.url
        })
    })
})
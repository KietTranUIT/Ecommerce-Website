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

    // xoa danh muc
    $('.btn-delete').click(function() {
        var userConfirmed = window.confirm("Deleting categories will likely affect products")
        if (!userConfirmed) {
            return
        }
        let id = this.dataset.categoryId
    
        let url = '/admin/categories/' + id
    
        fetch(url, {
            method: 'DELETE'
        })
        .then(response => response.json())
        .then(resJson => {
            if (resJson.status != true) {
                if(resJson.error_msg.includes('a foreign key constraint fails')) {
                    alert('The category cannot be deleted because there are related products')
                } else {
                    alert('Server Error!')
                }
            } else {
                alert('Delete successfull!')
                location.reload()
            }
        })
        })
    
    // Cap nhat danh muc
    $('.btn-edit').click(function() {
        let id = this.dataset.categoryId
        
        let url = '/admin/categories/update/' + id
        
        fetch(url, {
            method: 'GET'
        })
        .then(response => {
                window.location.href = response.url
            })
        })
})

function NewCategory() {
    fetch('/admin/categories/new', {
        method: 'GET'
    })
    .then(response => {
        window.location.href = response.url
    })
}
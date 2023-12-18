$(document).ready(function() {
    $('.btn-delete').click(function() {
    let id = this.dataset.productId

    let url = '/admin/products/' + id

    fetch(url, {
        method: 'DELETE'
    })
    .then(response => {
            console.log(response)
        })
    })

    $('.btn-edit').click(function() {
        let id = this.dataset.productId
    
        let url = '/admin/products/update/' + id
    
        fetch(url, {
            method: 'GET'
        })
        .then(response => {
                window.location.href = response.url
            })
    })

    $('.btn-detail').click(function() {
        let id = this.dataset.productId
    
        let url = '/admin/products/' + id
    
        fetch(url, {
            method: 'GET'
        })
        .then(response => {
                window.location.href = response.url
        })
    })
})

function NewProduct() {
fetch('/admin/products/new', {
    method: 'GET'
})
.then(response => {
    window.location.href = response.url
})
}
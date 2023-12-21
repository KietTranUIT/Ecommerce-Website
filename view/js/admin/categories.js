$(document).ready(function() {
    $('.btn-delete').click(function() {
    let id = this.dataset.productId

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
            alert('The category is deleted!')
        }
    })
    })

    $('.btn-edit').click(function() {
        let id = this.dataset.productId
    
        let url = '/admin/categories/update/' + id
    
        fetch(url, {
            method: 'GET'
        })
        .then(response => {
                window.location.href = response.url
            })
        })
})

function NewProduct() {
fetch('/admin/categories/new', {
    method: 'GET'
})
.then(response => {
    window.location.href = response.url
})
}
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
    var userConfirmed = window.confirm("Deleting products will likely affect inventory")
    if (!userConfirmed) {
        return
    }
    let id = this.dataset.productId

    let url = '/admin/products/' + id

    fetch(url, {
        method: 'DELETE'
    })
    .then(response => {
            if (response.status === 200) {
                alert('Product deleted')
                location.reload()
            } else {
                alert('Server error')
            }
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

    $('#form-filter').submit(function(event) {
        event.preventDefault()
    let category_id = $('#filter').val()
    fetch('/admin/products/filter?category_id=' + category_id, {
        method: 'GET'
    })
    .then(response => response.json())
    .then(data => {
        if (data.status != true) {
            alert("Filter fail!")
            return
        }

        let productsHTML = ''
        let products = data.data[0]
        if (products == null) {
            alert("Category don't product!")
            return;
        }
        products.forEach(product => {
            const html = `<tr>
            <td>${product.id}</td>
            <td>${product.name}</td>
            <td>
                <img src="${product.image}" alt="Error">
            </td>
            <td>
                <p style="width:500px;overflow:hidden;text-overflow:ellipsis">${product.description}</p></td>
            <td>${product.category_name}</td>
            <td>${formatCurrency(product.price)}</td>
            <td>
                <button type="button" data-product-id="${product.id}" class="btn btn-primary btn-edit">Edit</button>
                <button type="button" data-product-id="${product.id}" class="btn btn-outline-secondary btn-detail">Details</button>
                <button type="button" data-product-id="${product.id}" class="btn btn-danger btn-delete">Delete</button>
            </td>
        </tr>`

            productsHTML += html
        })
        let replace = `<tbody id="products-table">${productsHTML}</tbody>`
        $('#products-table').replaceWith(replace)
        $('.btn-delete').click(function() {
            var userConfirmed = window.confirm("Deleting products will likely affect inventory")
            if (!userConfirmed) {
                return
            }
            let id = this.dataset.productId
        
            let url = '/admin/products/' + id
        
            fetch(url, {
                method: 'DELETE'
            })
            .then(response => {
                    if (response.status === 200) {
                        alert('Product deleted')
                        location.reload()
                    } else {
                        alert('Server error')
                    }
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
    } )
})

function NewProduct() {
fetch('/admin/products/new', {
    method: 'GET'
})
.then(response => {
    window.location.href = response.url
})
}

function formatCurrency(number) {
    const formattedNumber = new Intl.NumberFormat('vi-VN', { style: 'currency', currency: 'VND' }).format(number);
    return formattedNumber;
}
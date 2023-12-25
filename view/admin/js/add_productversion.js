$(document).ready(function() {
    $('#form-insert').submit(function(event) {
        event.preventDefault()

        let p_id = $('#product_id').val()
        let size = $('#size').val()
        let quantity = $('#quantity').val()

        let jsonData = {
            id: "",
            p_id: parseInt(p_id),
            size_product: parseInt(size),
            inventory: quantity
        }
        let datajs = JSON.stringify(jsonData)

        let url = '/admin/products/' + p_id + '/products_version' 
        fetch(url, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: datajs
        })
        .then(response => {
            if(response.status === 200) {
                alert("Insert product version successfull")
                $('#size').val("")
                $('#quantity').val("")
            } else {
                alert("Insert product version failed")
            }
        })
    })
})
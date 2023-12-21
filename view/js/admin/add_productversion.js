$(document).ready(function() {
    $('.success').hide()
    $('.fail').hide()

    $('.success button').click(function() {
        $('.success').hide()
    })
    $('.fail button').click(function() {
        $('.fail').hide()
    })
    $('#create_product').submit(function(event) {
        event.preventDefault()

        let p_id = $('#p_id').val()
        let size = $('#size').val()
        let quantity = $('#inventory').val()

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
            $('.success').show()
            if(response.status === 200) {
                console.log("Them san pham thanh cong!")
                    $('.success').show()

            } else {
                console.log("Them san pham that bai!")
                    $('.fail').show()
            }
        })
    })
})
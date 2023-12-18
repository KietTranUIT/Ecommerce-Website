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

        let title = $('#title').val()
        let summary = $('#summary').val()
        let category_id = $('#list').val()
        let price = $('#price').val()

        if (category_id == '') {
            return
        }

        let jsonData = {
            name: title,
            description: summary,
            category_id: parseInt(category_id),
            description: summary,
            price: price
        }
        let datajs = JSON.stringify(jsonData)
        console.log(jsonData)

        fetch('/admin/products', {
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
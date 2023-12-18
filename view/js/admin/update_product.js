$(document).ready(function() {
    $('.success').hide()
    $('.fail').hide()

    $('.success button').click(function() {
        $('.success').hide()
    })
    $('.fail button').click(function() {
        $('.fail').hide()
    })

    $("#Save").submit(function(event) {
        event.preventDefault()

        let category_default = this.dataset.productId

        let id = $('#id').val()
        let name = $('#name').val()
        let description = $('#description').val()
        let category_id = $('#list').val()
        let price = $('#price').val()

        if (category_id == '') {
            return
        }

        if (category_id != category_default) {
            category_default = category_id
        }

        let data = {
            id: parseInt(id),
            name: name,
            description: description,
            category_id: parseInt(category_default),
            price: parseInt(price)
        }
        let dataJson = JSON.stringify(data)

        let url = '/admin/products/update/' + id

        fetch(url, {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            body: dataJson
        })
        .then(response => {
            if (response.status === 200) {
                console.log("Thay doi san pham thanh cong")
                    $('.success').show()
            } else {
                console.log("Thay doi san pham that bai")
                    $('.fail').show()
            }
        })
    })
})
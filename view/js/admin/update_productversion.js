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

        let id = $('#id').val()
        let pid = $('#p_id').val()
        let size = $('#size').val()
        let quantity = $('#inventory').val()

        let data = {
            id: parseInt(id),
            p_id: parseInt(pid),
            size_product: parseInt(size),
            inventory: quantity
        }
        let dataJson = JSON.stringify(data)
        console.log(dataJson)

        let url = '/admin/products_version/update/' + id

        fetch(url, {
            method: 'PUT',
            headers: {"Content-Type": "application/json"},
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
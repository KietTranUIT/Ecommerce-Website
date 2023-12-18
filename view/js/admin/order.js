$(document).ready(function() {
    $('.btn-detail').click(function() {
        let id = this.dataset.productId
    
        let url = '/admin/orders/' + id
    
        fetch(url, {
            method: 'GET'
        })
        .then(response => {
                window.location.href = response.url
        })
    })
})
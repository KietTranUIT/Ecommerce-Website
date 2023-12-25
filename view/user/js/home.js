
function GetProductsRequest() {
    let productsHTML = ''
    fetch('/products', {
        method: 'GET',
    })
    .then(response => response.json())
    .then(products => {
        products.data[0].forEach((product) => {
                    const html = `<div class="col-xl-3 col-lg-4 col-md-6 col-sm-6 col-12 product-card">
                    <div class="product-card-inner">
                        <div class="product-card-img">
                            <a href="/products/${product.id}" class="pro-img">
                                <img src="${product.image}">
                            </a>
                            <a class="wishlist" href="#" title="Wishlist" tabindex="0"><svg><use xlink:href="#wish"></use></svg></a>
                        </div>
                        <div class="product-card-content">
                            <h4><a href="#" tabindex="0">${product.name}</a></h4>
                            <p>${product.category_name}</p>
                            <div class="price"> 
                                <ins style="font-size:15px;">${product.price.toLocaleString('it-IT', {style : 'currency', currency : 'VND'})}</ins>
                            </div>
                            <a href="/products/${product.id}" class="add-cart-btn" tabindex="0" data-product-id=${product.id}>
                                <span>ADD TO CART</span>
                                <span class="atc-ic">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="9" height="8" viewBox="0 0 9 8" fill="none">
                                        <path fill-rule="evenodd" clip-rule="evenodd" d="M7.35342 5.2252H3.43344C2.90305 5.22535 2.44792 4.84736 2.35068 4.32595L1.84049 1.56215C1.8082 1.38557 1.65294 1.25825 1.47345 1.26118H0.621922C0.419212 1.26118 0.254883 1.09685 0.254883 0.894139C0.254883 0.691429 0.419212 0.5271 0.621922 0.5271H1.48079C2.01119 0.52695 2.46632 0.904941 2.56356 1.42635L3.07374 4.19015C3.10603 4.36673 3.2613 4.49405 3.44078 4.49112H7.35709C7.53657 4.49405 7.69184 4.36673 7.72413 4.19015L8.1866 1.69428C8.20641 1.58612 8.17667 1.47476 8.10558 1.39087C8.03448 1.30698 7.92951 1.25938 7.81956 1.26118H3.55824C3.35553 1.26118 3.1912 1.09685 3.1912 0.894139C3.1912 0.691429 3.35553 0.5271 3.55824 0.5271H7.81589C8.14332 0.527007 8.45381 0.672642 8.66308 0.924473C8.87235 1.1763 8.95868 1.50821 8.89865 1.83009L8.43619 4.32595C8.33895 4.84736 7.88381 5.22535 7.35342 5.2252ZM5.02645 6.69462C5.02645 6.08649 4.53347 5.59351 3.92534 5.59351C3.72263 5.59351 3.5583 5.75783 3.5583 5.96055C3.5583 6.16326 3.72263 6.32758 3.92534 6.32758C4.12805 6.32758 4.29238 6.49191 4.29238 6.69462C4.29238 6.89733 4.12805 7.06166 3.92534 7.06166C3.72263 7.06166 3.5583 6.89733 3.5583 6.69462C3.5583 6.49191 3.39397 6.32758 3.19126 6.32758C2.98855 6.32758 2.82422 6.49191 2.82422 6.69462C2.82422 7.30275 3.31721 7.79574 3.92534 7.79574C4.53347 7.79574 5.02645 7.30275 5.02645 6.69462ZM7.22865 7.4287C7.22865 7.22599 7.06433 7.06166 6.86162 7.06166C6.65891 7.06166 6.49458 6.89733 6.49458 6.69462C6.49458 6.49191 6.65891 6.32758 6.86162 6.32758C7.06433 6.32758 7.22865 6.49191 7.22865 6.69462C7.22865 6.89733 7.39298 7.06166 7.59569 7.06166C7.7984 7.06166 7.96273 6.89733 7.96273 6.69462C7.96273 6.08649 7.46975 5.59351 6.86162 5.59351C6.25349 5.59351 5.7605 6.08649 5.7605 6.69462C5.7605 7.30275 6.25349 7.79574 6.86162 7.79574C7.06433 7.79574 7.22865 7.63141 7.22865 7.4287Z" fill="white"></path>
                                    </svg>
                                </span>
                            </a>
                        </div>
                    </div>
                </div>`
                productsHTML += html
                })
                let result = `<div class="row" id="list-row-products">${productsHTML}</div>`
             $('#list-row-products').replaceWith(result)

    })
}

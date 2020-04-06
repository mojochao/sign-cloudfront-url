# sign-cloudfront-url

This application outputs a signed URL for a CloudFront URL.

# Usage

Clone the repository:

    $ git clone https://github.com/mojochao/sign-cloudfront-url.git
    
Build the executable:
    
    $ cd sign-cloudfront-url
    $ go build -o sign-cloudfront-url .

Run the executable:

    $ ./sign-cloudfront-url https://dijb73lf0jsnl.cloudfront.net/index.html SOMEKEYPAIRID some_private_key.pem
    reading private key from some_private_key.pem ...
    signing url https://dijb73lf0jsnl.cloudfront.net/index.html with keypair APKAJMQH5OQEYWRFQBVQ and private key ...
    
    https://dijb73lf0jsnl.cloudfront.net/index.html?Policy=eyJTdGF0ZW1lbnQiOlt7
    IlJlc291cmNlIjoiaHR0cHM6Ly9kaXhrNzNsZjBqbnNsLmNsb3VkZnJvbnQubmV0L2luZGV4Lmh
    0bWwiLCJDb25kaXRpb24iOnsiRGF0ZUdyZWF0ZXJUaGFuIjp7IkFXUzpFcG9jaFRpbWUiOjE1OD
    YwMzgwMDd9LCJEYXRlTGVzc1RoYW4iOnsiQVdTOkVwb2NoVGltZSI6MTU4NjEyNDQwN319fV19&
    Signature=dwr0zEDvOrB0rpyDYGko1RDoHKzwj1uUmHMv4uZcQQYIYTPNVuh~xkVUpJmR0~UJe
    l9kCdntrJRpVeCiFYU0PqmBb6Kx65-UaUtn2~bT1L1U30iDrm5GYFs2mc8dJvagCZfTmzHnmo9U
    HeycqXQ53PfOrW0EPrO~LY7Pu8mNtPcQddBfGUjq65oC4G3pLz3zAwIDPyqApBojFUpg7rBf65P
    NjWSprZHIH9s5CQnfydd61HW14ppq3mDPDbC7EwlI4jw53~97fXdD3rhaF5hKBMfN5SgAg80uiV
    6iu8Oa5yzdQcXV2L2lnZ1HEnQgcfGLobu3iDcBPTa9yM4fYLN9TbL0QErXpkjPpaYItgnhlvDXc
    TTeP-BXkz4IUktcwuPWB5TvP3ug8Lwil-0Ni0diJGU~L7mBI11xd2ceCEJZId2YMmHy2b7nXsO6
    A8PzkOMUF31Ej7-R8tS4~StKttK1IwYoyrmXLeg2T31bkmH-6ALAQXtXtumZ8ra4lIMpZAbyuAT
    Z7AYEFtEXkDNaXweNWDNwHdAQ0tSdgfzvKFAImmu5HuqHCrIqwKAOHCj~8h0vZzHckHHuDS6WIn
    jBNIZdxBVbARUrFQgysbSLAtjELLe5Ln6OuoPUhA-~n91vsaOw4S0UsBY-UyJQBY9Cf2Lg-vjKe
    wouS7ZyxvRbpr~bZEg_&Key-Pair-Id=SOMEKEYPAIRID
    
Try to open the signed url in bottom of output.

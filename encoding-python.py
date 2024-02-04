s1 = "ABC"
s2 = "!#5"
s3 = "😁😁😁"
li = [s1, s2, s3]
#UTF-8 encoding of 😁 takes up four bytes
#But since the Python len() method returns the number of Unicode code points,
#we get 3 for each of the above three strings
if __name__ == "__main__":
    for s in li:
        print(len(s))
    #print(len(s1)) //3
    #print(len(s2)) //3
    #print(len(s3)) //3
    #Get byte length for each of the strings
    print('*************')
    for s in li:
        encoded_string = s.encode('utf-8')
        print(len(encoded_string))

    print('*************')
    for s in li:
        encoded_string = s.encode('utf-16')
        print(len(encoded_string))

    '''
    the len() function in Python returns the number of characters in a string based on Unicode representation,
    and it does not directly consider the byte size of individual characters in memory. If you need to work with byte sizes, you should encode the string and then use the len() function on the resulting byte sequence

    my_string = "Hello, World!"
encoded_string = my_string.encode('utf-8')
byte_length = len(encoded_string)
print(f"The byte length of the string is: {byte_length}")

    
    '''
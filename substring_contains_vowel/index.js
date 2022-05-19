function isVowel(x)
{
    // Function to check whether a character is
    // vowel or not
    return (x == 'a' || x == 'e' || x == 'i'
            || x == 'o' || x == 'u');       
}
 
function findSubstring(str)
{
    let hash = new Set();
        // To store vowels

    // Outer loop picks starting character and
    // inner loop picks ending character.
    let n = str.length;
    for (let i = 0; i < n; i++) {
        for (let j = i; j < n; j++) {

            // If current character is not vowel,
            // then no more result substrings
            // possible starting from str[i].
            if (isVowel(str[j]) == false)
                break;

            // If vowel, then we insert it in hash
            hash.add(str[j]);
            console.log(hash)

            // If all vowels are present in current
            // substring
            if (hash.size == 5)
                console.log(str.substring(i, j + 1) + " ");
        }
        hash.clear();
    }
}
 
// Driver code
let str = "aeoibsddaeiouudb";
findSubstring(str);
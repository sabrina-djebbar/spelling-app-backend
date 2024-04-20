package app

import ("context"
"github.com/sabrina-djebbar/spelling-app-backend/lib/id")

func (a *app) CreateWord(ctx context.Context, spelling string) error {
	word := &models.Word{
		id: id.Generate("word")
		spelling: spelling
		difficuly: CalculateDifficulty(spelling)
	}
}

/**  1. **Word Length**: Longer words are generally more difficult to spell than shorter ones. Calculate the length of the word.

2. **Commonality of Sounds**: Break down the word into phonemes (individual sounds) and consider how common or uncommon these phonemes are in the English language. Use a phoneme database or library to determine the frequency of each phoneme.

3. **Irregularities**: Check if the word follows common spelling rules or if it contains irregular spellings. Words with irregular spellings are generally considered more difficult.

4. **Syllable Complexity**: Analyze the syllable structure of the word. Words with complex syllable structures or unusual syllable patterns may be more difficult to spell.

5. **Word Frequency**: Consider the frequency of the word in everyday language usage. Less commonly used words may be more difficult for students to remember.

6. **Contextual Difficulty**: Consider the context in which the word is used. Some words may be more difficult to spell because they have multiple meanings or are used in specific contexts.

7. **Homophones and Homographs**: Check if the word is a homophone (sounds the same as another word but has a different meaning and spelling) or a homograph (spelled the same as another word but has a different meaning and possibly pronunciation). Homophones and homographs can add complexity to spelling.

8. **Word Origin and Etymology**: Consider the origin and etymology of the word. Words with roots in other languages or historical influences may have more complex spellings.

9. **Rule Exceptions**: Identify any exceptions to common spelling rules that apply to the word.

10. **Contextual Usage**: Analyze how the word is typically used in sentences and paragraphs. Words that are commonly misspelled in context may be considered more difficult.

11. **Compute Difficulty Score**: Assign a numerical score to each of the above factors based on their impact on the difficulty of the word. Combine these scores to calculate an overall difficulty score for the word.

12. **Ranking**: Rank the words based on their difficulty scores, with higher scores indicating greater difficulty.

This algorithm provides a general framework for assessing the difficulty of spelling words according to the UK government's national curriculum. Depending on your specific needs and requirements, you may need to adjust or expand upon this algorithm.
*/

func CalculateDifficulty(spelling string)int{

}
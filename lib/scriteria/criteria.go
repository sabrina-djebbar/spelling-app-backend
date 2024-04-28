package scriteria

func GetSpellingPatterns() []Pattern {
	patterns := []Pattern{
		{5.0, "Year 1", "ff, ll, ss, zz, ck", `(ff|ll|ss|zz|ck)`},
		{4.0, "Year 1", "n before k", `n[Kk]`},
		{4.0, "Year 1", "-tch", `(tch)`},
		{3.0, "Year 1", "/v/ sound at the end of words", `\b\w*v\b`},
		{4.0, "Year 1", "Adding s and es to words", `\b\w+s\b|\b\w+es\b`},
		{4.0, "Year 1", "Adding the endings –ing, –ed and –er to verbs", `\b\w+(ing|ed|er)\b`},
		{3.0, "Year 1", "Adding –er and –est to adjectives", `\b\w+(er|est)\b`},
		{5.0, "Year 1", "Vowel digraphs ai, oi", `\b(ai|oi)\b`},
		{5.0, "Year 1", "Vowel digraphs ay, oy", `\b(ay|oy)\b`},
		{6.0, "Year 1", "Vowel digraphs a-e, e-e, i-e, o-e, u-e", `\b(a-e|e-e|i-e|o-e|u-e)\b`},
		{3.0, "Year 1", "Vowel digraphs ar", `\b(ar)\b`},
		{4.0, "Year 1", "Vowel digraphs ee", `\b(ee)\b`},
		{4.0, "Year 1", "Vowel digraphs ea", `\b(ea)\b`},
		{4.0, "Year 1", "Vowel digraphs er", `\b(er)\b`},
		{4.0, "Year 1", "Vowel digraphs or", `\b(or)\b`},
		{3.0, "Year 1", "Vowel digraphs ow", `\b(ow)\b`},
		{3.0, "Year 1", "Vowel digraphs ou", `\b(ou)\b`},
		{4.0, "Year 1", "Vowel digraphs oi", `\b(oi)\b`},
		{3.0, "Year 1", "Vowel digraphs oy", `\b(oy)\b`},
		{4.0, "Year 1", "Vowel digraphs ue", `\b(ue)\b`},
		{4.0, "Year 1", "Vowel digraphs ew", `\b(ew)\b`},
		{3.0, "Year 1", "Vowel digraphs ue", `\b(ue)\b`},
		{4.0, "Year 1", "Vowel digraphs igh", `\b(igh)\b`},
		{4.0, "Year 1", "Vowel digraphs ie", `\b(ie)\b`},
		{3.0, "Year 1", "Vowel digraphs ir", `\b(ir)\b`},
		{3.0, "Year 1", "Vowel digraphs ur", `\b(ur)\b`},
		{3.0, "Year 1", "Vowel digraphs ay", `\b(ay)\b`},
		{3.0, "Year 1", "Vowel digraphs ow", `\b(ow)\b`},
		{4.0, "Year 1", "Words ending -y", `\b\w+y\b`},
		{4.0, "Year 1", "New consonant spellings ph and wh", `\b(wh|ph)\b`},
		{4.0, "Year 1", "Using k for the /k/ sound", `\b\w*k\w*\b`},
		{4.0, "Year 1", "Adding the prefix –un", `\bun\w*\b`},
		{4.0, "Year 1", "Compound words", `\b\w+\s\w+\b`},
		{3.0, "Year 1", "Common exception words", `\b(the|a|do|to|today|of|said|says|are|were|was|is|his|has|I|you|your|they|be|he|me|she|we|no|go|so|by|my|here|there|where|love|come|some|one|once|ask|friend|school|put|push|pull|full|house|our)\b`},
		{5.0, "Year 2", "The /dʒ/ sound spelt as ge and dge at the end of words", `\b(ge|dge)\b`},
		{4.0, "Year 2", "The /dʒ/ sound spelt as g elsewhere in words before e, i and y", `\bg[eiy]\b`},
		{4.0, "Year 2", "The /s/ sound spelt c before e, i and y", `\bc[eiy]\b`},
		{4.0, "Year 2", "The /n/ sound spelt kn and (less often) gn at the beginning of words", `\b(kn|gn)\w*\b`},
		{4.0, "Year 2", "The /r/ sound spelt wr at the beginning of words", `\bwr\w*\b`},
		{3.0, "Year 2", "The /l/ or /əl/ sound spelt –le at the end of words", `\b\w*le\b`},
		{3.0, "Year 2", "The /l/ or /əl/ sound spelt –el at the end of words", `\b\w*el\b`},
		{3.0, "Year 2", "The /l/ or /əl/ sound spelt –al at the end of words", `\b\w*al\b`},
		{4.0, "Year 2", "Words ending –il", `\b\w*il\b`},
		{3.0, "Year 2", "The /aɪ/ sound spelt –y at the end of words", `\b\w*y\b`},
		{4.0, "Year 2", "Adding –es to nouns and verbs ending in –y", `\b\w+ies\b`},
		{4.0, "Year 2", "Adding –ed, –ing, –er and –est to a root word ending in –y with a consonant before it", `\b\w+(ied|ying|ier|iest)\b`},
		{3.0, "Year 2", "Adding the endings –ing, –ed, –er, –est and –y to words ending in –e with a consonant before it", `\b\w*(ing|ed|er|est|y)\b`},
		{3.0, "Year 2", "Adding –ing, –ed, –er, –est and –y to words of one syllable ending in a single consonant letter after a single vowel letter", `\b\w*(ing|ed|er|est|y)\b`},
		{4.0, "Year 2", "The /ɔ:/ sound spelt a before l and ll", `\b\w*al\b`},
		{3.0, "Year 2", "The /ʌ/ sound spelt o", `\bo\w*\b`},
		{4.0, "Year 2", "The /i:/ sound spelt –ey", `\b\w*ey\b`},
		{4.0, "Year 2", "The /ɒ/ sound spelt a after w and qu", `\b(wa|qua)\w*\b`},
		{3.0, "Year 2", "The /ɜ:/ sound spelt or after w", `\b(wor)\w*\b`},
		{3.0, "Year 2", "The /ɔ:/ sound spelt ar after w", `\b(war)\w*\b`},
		{3.0, "Year 2", "The /ʒ/ sound spelt s", `\bs\w*\b`},
		{4.0, "Year 2", "The suffixes –ment, –ness, –ful, –less and –ly", `\b(\w*ment|\w*ness|\w*ful|\w*less|\w*ly)\b`},
		{3.0, "Year 2", "Contractions", `\b(can't|didn't|hasn't|couldn't|it's|I'll)\b`},
		{3.0, "Year 2", "The possessive apostrophe (singular nouns)", `\b(\w*\'s)\b`},
		{4.0, "Year 2", "Words ending in –tion", `\b(\w*tion)\b`},
		{3.0, "Year 2", "Homophones and near-homophones", `\b(there|their|they're|here|hear|quite|quiet|see|sea|bare|bear|one|won|sun|son|to|too|two|be|bee|blue|blew|night|knight)\b`},
		{3.0, "Year 2", "Common exception words", `\b(the|I|are|were|was|you|your|they|be|me|she|we|he|no|go|so|by|my|here|there|where|come|some|one|once|ask|friend|school|put|push|pull|full|our)\b`},
	}
	return patterns
}
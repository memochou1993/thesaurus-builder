const app = document.querySelector('#app');
const spinner = document.querySelector('#spinner');
const title = document.querySelector('#title');
const root = document.querySelector('#root');
const subjectTemplate = document.querySelector('[data-subject-template]');
const termTemplate = document.querySelector('[data-term-template]');
const noteTemplate = document.querySelector('[data-note-template]');

/**
 * @param {HTMLElement} target
 * @param {Object} prop
 * @param {Object} prop.subject
 * @param {Object} prop.subject.terms
 * @param {Array}  prop.subject.terms[].text
 * @param {Array}  prop.subject.terms[].preferred
 * @param {Object} prop.subject.notes
 * @param {string} prop.subject.notes[].text
 * @param {Array}  prop.children
 */
const render = (target, prop) => {
  const [subject] = subjectTemplate.content.cloneNode(true).children;
  const [preferredTerm, cptNotes, notes, cptTerms, terms, cptNarrowerTerms, narrowerTerms] = subject.children;
  if (!prop.subject.notes?.length) {
    cptNotes.classList.add('hidden');
    notes.classList.add('hidden');
  }
  if (prop.subject.terms.length === 1) {
    cptTerms.classList.add('hidden');
    terms.classList.add('hidden');
  }
  if (!prop?.children?.length) {
    cptNarrowerTerms.classList.add('hidden');
    narrowerTerms.classList.add('hidden');
    preferredTerm.classList.remove('pointer');
  }
  prop.subject.terms.forEach((item) => {
    if (item.preferred) {
      preferredTerm.textContent = item.text;
    }
    const [term] = termTemplate.content.cloneNode(true).children;
    const [text] = term.children;
    text.textContent = item.preferred ? `${item.text} (preferred)` : item.text;
    terms.appendChild(term);
  });
  prop.subject.notes?.forEach((item) => {
    const [note] = noteTemplate.content.cloneNode(true).children;
    const [text] = note.children;
    text.textContent = item.text;
    notes.append(note);
  });
  target.appendChild(subject);
  setTimeout(() => prop?.children?.forEach((item) => render(narrowerTerms, item)), 0);
};

const toggleSpinner = async (delay = 0) => {
  await new Promise((res) => setTimeout(() => res(), delay));
  document.documentElement.classList.toggle('full-height');
  spinner.classList.toggle('hidden');
  app.classList.toggle('hidden');
};

(async () => {
  await toggleSpinner();
  const data = await fetch('data.json').then((r) => r.json());
  title.textContent = data.title;
  render(root, data.root);
  await toggleSpinner(1000);
})();

root.addEventListener('click', (e) => {
  if (e.target.classList.contains('pointer')) {
    e.target.parentElement.querySelector('.caption-narrower-terms').classList.toggle('hidden');
    e.target.parentElement.querySelector('.narrower-terms').classList.toggle('hidden');
    e.target.classList.toggle('preferred-term-clicked');
  }
});

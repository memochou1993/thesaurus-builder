/**
 * @param {HTMLElement} target
 * @param {Object} prop
 * @param {Object} prop.subject
 * @param {Object} prop.subject.term
 * @param {Array}  prop.subject.term.preferredTerms
 * @param {string} prop.subject.term.preferredTerms[].termText
 * @param {Object} prop.subject.note
 * @param {Array}  prop.subject.note.descriptiveNotes
 * @param {string} prop.subject.note.descriptiveNotes[].noteText
 * @param {Array}  prop.children
 */
const render = (target, prop) => {
  const template = document.querySelector('[data-subject-template]');
  const noteTemplate = document.querySelector('[data-descriptive-note-template]');
  const [subject] = template.content.cloneNode(true).children;
  const [term, note, children] = subject.children;
  const [descriptiveNotes] = note.children;
  prop.subject.term.preferredTerms.forEach((v) => {
    term.textContent = v.termText;
    term.classList.add(prop?.children?.length ? 'term-expandable' : 'term-expanded');
  });
  prop.subject.note?.descriptiveNotes?.forEach((v) => {
    const [descriptiveNote] = noteTemplate.content.cloneNode(true).children;
    const [noteText] = descriptiveNote.children;
    noteText.textContent = v.noteText;
    descriptiveNotes.append(descriptiveNote);
  });
  prop.children?.forEach((c) => render(children, c));
  target.appendChild(subject);
};

(async () => {
  const data = await fetch('data.json').then((r) => r.json());
  const title = document.querySelector('#title');
  const root = document.querySelector('#root');
  title.textContent = data.title;
  root.addEventListener('click', (e) => {
    if (e.target.classList.contains('term-expandable')) {
      e.target.parentElement.querySelector('.children').classList.toggle('active');
      e.target.classList.toggle('term-expanded');
    }
  });
  render(root, data.root);
})();

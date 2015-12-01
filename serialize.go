package libxml2

import (
	"github.com/lestrrat/go-libxml2/dom"
	"github.com/lestrrat/go-libxml2/clib"
	"github.com/lestrrat/go-libxml2/node"
)

// Serialize produces serialization of the document, canonicalized.
func (s C14NSerialize) Serialize(n node.Node) (string, error) {
	/*
	 * Below document is taken from libxml2 directly. Pay special attention
	 * to the required settings when parsing the document to be canonicalized.
	 *
	 * ---
	 * Canonical form of an XML document could be created if and only if
	 *  a) default attributes (if any) are added to all nodes
	 *  b) all character and parsed entity references are resolved
	 * In order to achive this in libxml2 the document MUST be loaded with
	 * following global setings:
	 *
	 *    xmlLoadExtDtdDefaultValue = XML_DETECT_IDS | XML_COMPLETE_ATTRS;
	 *    xmlSubstituteEntitiesDefault(1);
	 *
	 * or corresponding parser context setting:
	 *    xmlParserCtxtPtr ctxt;
	 *
	 *    ...
	 *    ctxt->loadsubset = XML_DETECT_IDS | XML_COMPLETE_ATTRS;
	 *    ctxt->replaceEntities = 1;
	 *    ...
	 * ---
	 *
	 * In go-libxml2, this translates to:
	 *
	 *    options = XMLParserDTDLoad | XMLParserDTDAttr | XMLParserNoEnt
	 *
	 */
	switch n.(type) {
	case *dom.Document:
	default:
		return "", ErrInvalidNodeType
	}

	return clib.XMLC14NDocDumpMemory(n, int(s.Mode), s.WithComments)
}

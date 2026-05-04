from typing import TypedDict, List, Optional

class Revision(TypedDict):
    data: str
    status: str  # "x" para feito, " " para pendente

class Reference(TypedDict):
    description: str
    source: str

class NoteMetadata(TypedDict):
    filename: str
    relative_path: str
    tema: str
    subtema: Optional[str]
    revisoes: List[Revision]
    tags: List[str]
    references: List[Reference]

class ProcessedNote(NoteMetadata):
    dias_atraso: int
    status_revisao: str  # "ATRASADA", "HOJE", "FUTURA"

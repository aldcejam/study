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
    activity: Optional[str]
    updatedAt: str # Data da última alteração via Git (ISO)

class ProcessedNote(NoteMetadata):
    id: str # Identificador único: caminho + data_revisao
    dias_atraso: int
    status_revisao: str  # "ATRASADA", "HOJE", "FUTURA"
    last_notified_at: Optional[str] # Data do último alerta enviado (ISO)
